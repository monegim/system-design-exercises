package main

import "helmify/pkg/config"

const helpText = `Helmify parses kubernetes resources from std.in and converts it to a Helm chart.
Example 1: 'kustomize build <kustomize_dir> | helmify mychart' 
  - will create 'mychart' directory with Helm chart from kustomize output.
Example 2: 'cat my-app.yaml | helmify mychart' 
  - will create 'mychart' directory with Helm chart from yaml file.
Example 3: 'awk 'FNR==1 && NR!=1  {print "---"}{print}' /my_directory/*.yaml | helmify mychart' 
  - will create 'mychart' directory with Helm chart from all yaml files in my_directory directory.
Usage:
  helmify [flags] CHART_NAME  -  CHART_NAME is optional. Default is 'chart'. Can be a directory, e.g. 'deploy/charts/mychart'.
Flags:
`

func ReadFlags() config.Config {
	result := config.Config{}

	return result
}
