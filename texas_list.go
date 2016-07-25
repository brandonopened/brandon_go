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
  glog.Infoln("Texas Standards")
token,err := opened.GetToken("","","","")
  query_params:=make(map[string]string)
  query_params["descriptive"]="fractions"
  query_params["standard_group"]="21"
  query_params["grades_range"]="K-12"
  results,err:=opened.SearchResources(query_params,token)
  if err!=nil {
    glog.Errorf("Error from SearchResources %+v",err)
  }
  glog.V(1).Infof("%d results returned",len(results.Resources))
  for _,resource:= range results.Resources {
    //I want the output to simply be "standard id, count number of resources"
    glog.V(2).Infof("Resource %d (%s)",resource.Id,resource.Title)
  }
}
