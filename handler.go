package main

import (
	"encoding/json"
	"log"
	"main/addressBook"
	"net/http"
)

type (
	request struct {
		ReqType string `json:"req_type"`
		Data    []struct {
			Item string `json:"item"`
		} `json:"data"`
	}
	response struct {
		ResType string `json:"res_type"`
		Result  string `json:"result"`
		Data    string `json:"data"`
	}
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	var (
		req   request  //request struct
		resp  response //response struct
		err   error
		list  = make([]string, 0) //accumulated address list from request
		jResp = make([]byte, 0)   //marshalled response
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		badRequest(w, err)
		return
	}
	for _, v := range req.Data {
		if v.Item == "" {
			continue
		} else {
			list = append(list, v.Item)
		}
	}
	resp.ResType = req.ReqType
	if resp.Data, err = new(addressBook.Book).Format(list); err != nil {
		badRequest(w, err)
		return
	}

	if jResp, err = json.Marshal(resp); err != nil {
		serverError(w, err)
		return
	}
	_, _ = w.Write(jResp)
	return
}

func badRequest(w http.ResponseWriter, err error) {
	var resp response
	log.Println("[WARN] bad body provided: ", err)
	resp.ResType = ""
	resp.Result = "fail"
	resp.Data = err.Error()
	jResp, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write(jResp)
	return
}

func serverError(w http.ResponseWriter, err error) {
	var resp response
	log.Println("[ERR] unexpected error: ", err)
	resp.ResType = ""
	resp.Result = "fail"
	resp.Data = err.Error()
	jResp, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(jResp)
	return
}
