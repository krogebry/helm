# Deploy a simple helm chart from a local directory

This is used for testing helm charts before being published to Artifactory

#### Example

```yaml
      - id: local-deploy
        uses: Innersource/build-infra-actions/helm-deploy-local@v0.1.50
        with:
          stack_id: ${{ inputs.helm_chart_name }}-${{ env.SHORT_SHA }}
          helm_chart_name: ${{ inputs.helm_chart_name }}
          helm_chart_path: ${{ inputs.helm_chart_path }}/${{ inputs.helm_chart_name }}
```

# Inputs
|Input Name|Description|Required|Default|Type|
|---|---|---|---|---|
|environment|Which environment to deploy into|True| |choice|
|image_tag|Container version to deploy|True| |string|
|helm_chart_name|Helm chart name|True| |string|
|helm_chart_path|Helm chart path|True| |string|
|extra_sets|Extra --set flags|False||string|
|helm_values_path|Where to find the helm chart value overrides per environment|False|.|string|

