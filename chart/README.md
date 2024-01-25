## Caveat of creating a custom chart for devtron
1. Create a .image_descriptor_template.json and populate it with the contents shown in the dock
2. Create a app-values.yaml file and populate it with the contents of values.yaml
3. Create a release-values.yaml and populate it with the values from https://github.com/devtron-labs/devtron/blob/main/scripts/devtron-reference-helm-charts/deployment-chart_4-18-0/release-values.yaml
4. Package the app with `helm package ./chart`

