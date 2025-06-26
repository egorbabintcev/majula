package web

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	service Service
}

func newHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) handleGetWhoAmI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, _ := json.Marshal(GetWhoAmIRes{
		Username: "system",
	})

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *handler) handleGetPkg(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "package")

	p, err := h.service.GetPkg(name)

	if err != nil {
		respondWErr(w, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	res, _ := json.Marshal(GetPkgRes{
		Name:     p.Name,
		Versions: p.Versions,
		Tags:     p.Tags,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *handler) handlePutPkg(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		respondWErr(w, http.StatusInternalServerError)
		return
	}

	req := PutPkgReq{}

	err = json.Unmarshal(body, &req)

	if err != nil {
		respondWErr(w, http.StatusInternalServerError)
		return
	}

	name := chi.URLParam(r, "package")

	if name != req.Name {
		respondWErr(w, http.StatusInternalServerError)
		return
	}

	if len(req.Versions) == 0 {
		respondWErr(w, http.StatusInternalServerError)
		return
	}

	version := ""

	for k := range req.Versions {
		version = k
		break
	}

	manifest := req.Versions[version]

	tags := make([]string, 0)

	for k, v := range req.Tags {
		if v == version {
			tags = append(tags, k)
		}
	}

	tarName := fmt.Sprintf("%s-%s.tgz", name, version)

	if _, exist := req.Attachments[tarName]; !exist {
		respondWErr(w, http.StatusInternalServerError)
		return
	}

	attachments := req.Attachments[tarName]

	if attachments.ContentType != "application/octet-stream" {
		respondWErr(w, http.StatusInternalServerError)
		return
	}

	tarball, err := base64.StdEncoding.DecodeString(attachments.Data)

	if err != nil {
		respondWErr(w, http.StatusInternalServerError)
		return
	}

	if len(tarball) != attachments.Length {
		respondWErr(w, http.StatusInternalServerError)
		return
	}

	err = h.service.PublishPkg(name, version, tags, manifest, tarball)

	if err != nil {
		respondWErr(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *handler) handleGetTarball(w http.ResponseWriter, r *http.Request) {
	t := chi.URLParam(r, "tarball")

	res, err := h.service.GetTarball(t)

	if err != nil {
		respondWErr(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(res.Content)
}
