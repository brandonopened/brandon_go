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
  query_params["standard_group"]="1"
  query_params["grades_range"]="K"
  query_params["resource_type"]="video"
  results,err:=opened.SearchResources(query_params,token)
  if err!=nil {
    glog.Errorf("Error from SearchResources %+v",err)
  }
  glog.V(1).Infof("%d results returned",len(results.Resources))
  for _,resource:= range results.Resources {
  // Find by resource type
    glog.V(2).Infof("Resource Title: %s (%d)",standard_group,resource.Count)
  }
}
