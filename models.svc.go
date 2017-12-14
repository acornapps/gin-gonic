package main

import "errors"

type service struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var serviceList = []service{
	service{ID: 1, Title: "Apps Hello world Json", Content: "Apps Hello World Contest~"},
	service{ID: 2, Title: "Apps Sample Rest Json", Content: "Apps Sample Contest! ~"},
}

func getAllServiceList() []service {
	return serviceList
}

func getServiceByID(id int) (*service, error) {
	for _, a := range serviceList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Service not found")
}