package main

import (
  "flag"
   "github.com/golang/glog"
   "github.com/openedinc/opened-go"
)

// init parses the flag options
func init() {
  flag.Parse()
}

func main() {
  glog.Infoln("My OpenEd Resource Find Script")
token,err := opened.GetToken("","","","")
  query_params:=make(map[string]string)
  query_params["standard"]="K.G.2"
  query_params["descriptive"]=""
  query_params["grades_range"]="K-12"
  results,err:=opened.SearchResources(query_params,token)
  if err!=nil {
    glog.Errorf("Error from SearchResources %+v",err)
  }
  glog.V(1).Infof("%d results returned",len(results.Resources))
  for _,resource:= range results.Resources {
    glog.V(2).Infof("Resource Title: %s (%d)",resource.Title,resource.Id)
  }
}
