# Sample Steampipe plugin

This is a starter [Steampipe](https://steampipe.io) plugin that you can customize to your needs.

It can be used to bootstrap any plugin, since it doesn't yet include dependencies for any SDK/external service.

## Usage

- [ ] Search for and replace all appearances of `jreyesr` with your Github username. Go mixes repository location and package name, so use wherever the code will be uploaded
- [ ] Check if there is a Go package for whatever service you will integrate. If so, run `go get PACKAGE_NAME`
- [ ] Search for and replace all appearances of `samplerest` with a short name for the target service. For example, `github`, `jira`, `aws` (these already exist)
- [ ] Edit the `config/samplerest.spc` to contain the credentials that are required to connect to the target service (URL, port, username, password, API keys, subdomain, account ID, whatever).Leave the `plugin` key, it is required by Steampipe
- [ ] Rename the `config/samplerest.spc` file to match the target service
- [ ] On the `samplerest/config.go` file:
    - [ ] Edit the `SampleRESTConfig` struct: rename it, edit its fields to match whatever you configured in the `config/*.spc` file. Ensure that the `` `cty:"something"` `` annotations match the keys in the `.spc` file
    - [ ] Edit the `ConfigSchema` variable to match
    - [ ] Edit the `String()` function to print an appropriate representation of your credentials. Endure that secrets are not leaked!
- [ ]  On the `samplerest/plugin.go` file:
    - [ ] Change the `Name`
- [ ] For each table that you want to expose via SQL:
    - [ ] Copy the `samplerest/table_samplerest_one_model.go` file
    - [ ] Rename it to describe the service (instead of `samplerest`) and the entity listed (instead of `one_model`). For example, `table_github_repository`
    - [ ] Change the `Name` and `Description`
    - [ ] If your model doesn't support searching to return a subset of items, delete the `List.KeyColumns` field, and any other places marked with `Delete if your API doesn't suport searching over all instances`
    - [ ] Add/edit all column names types and descriptions in `Columns` to match whatever is exposed by the API. The `Name` field will be seen by SQL, and the `Transform` field is used to match the objects that are returned by the `List` and `Get` functions
    - [ ] Rename the `OneModel` struct, and edit it to match the data exposed by the API. The field names should match with the names passed to the `Columns.Transform` configs above
    - [ ] Edit the `listOneModel` function to contact the API and get the results. You have available the `config` var, which holds API credentials, and possibly the `realQueryString` and/or `realQueryJson` variables, for filtering
    - [ ] Complete the `listOneModel` function to make it return all data returned by the API
    - [ ] Edit the `getOneModel` function to contact the API and get a single result. You have available the `config` var, which holds API credentials, and the `id` var, which holds the ID of the single object
    - [ ] Complete the `getOneModel` function to make it return the data of a single item
    - [ ] Rename the `listOneModel` and `getOneModel` functions to something that matches the actual objects. For example, `listRepository` and `getRepository` for the file `table_github_repository.go`
- [ ]  On the `samplerest/plugin.go` file:
    - [ ] Register your new tables in the `TableMap` field
    - [ ] Delete the sample table
- [ ] Rename the `samplerest` directory, where the code lives, to something that matches you service

## Testing

Run `make`, then run `steampipe query`. Run `.inspect` inside of it to ensure that the plugin is loaded.

Alternatively, run `go build -o ~/.steampipe/plugins/hub.steampipe.io/plugins/YOUR_USERNAME/SERVICENAME@latest/steampipe-plugin-SERVICENAME.plugin *.go`, replacing `YOUR_USERNAME` and `SERVICENAME`.