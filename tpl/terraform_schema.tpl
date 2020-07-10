{{ define "TerraformSchema" }}
  <div class="container">
    <div class="row justify-content-start mt-3 pl-2">
      <div class="col-8-md">
        <h2>Terraform Schema Tool</h2>
        <p>Easily create <a href="https://www.terraform.io/docs/extend/schemas/index.html" target="_blank">schema</a> definitions for your Terraform provider.</p>
      </div>
    </div>

    <div class="row justify-content-around mt-4">
      <div class="col-6-md">
        <div class="mr-5">
          <form>
            <div id="alerts">

            </div>
            <div class="form-group">
              <label for="name">Name</label>
              <input type="text" class="form-control-sm" id="name" aria-describedby="name"
                     placeholder="Name" required>
            </div>

            <div class="form-group">
              <label for="type">Type</label>
              <select class="form-control-sm" id="type">
                <option>Bool</option>
                <option>Int</option>
                <option>Float</option>
                <option>String</option>
              </select>
            </div>

            <div id="enable-default-form" class="form-check">
              <input class="form-check-input" type="checkbox" value="" id="enable-default">
              <label class="form-check-label" for="enable-default">
                Enable Default Value
              </label>
            </div>

            <div id="default-form" class="form-group" style="display: none;">
              <label for="default-value">Default Value</label>
              <select class="form-control-sm default-input" id="default-value-bool">
                <option>true</option>
                <option>false</option>
              </select>
              <input type="number" class="form-control-sm default-input" id="default-value-number" style="display: none;">
              <input type="text" class="form-control-sm default-input" id="default-value-string" style="display: none;">
            </div>

            <div class="form-group">
              <label for="constraint">Constraint</label>
              <select class="form-control-sm" id="constraint">
                <option>Required</option>
                <option>Optional</option>
                <option>Computed</option>
              </select>
            </div>

            <div class="form-check">
              <input class="form-check-input" type="checkbox" value="" id="force-new">
              <label class="form-check-label" for="force-new">
                Force New
              </label>
            </div>

            <button type="button" id="new-attribute-btn" class="btn btn-primary">Add New</button>
          </form>
        </div>
      </div>

      <div class="col-6-md">
        <ul id="attribute-box" class="list-group">
        </ul>
      </div>
    </div>

    <div class="row justify-content-center mt-4">
      <div class="col-8-md border border-dark rounded">
        <pre class="p-1"><code id="tf-code">
Schema: map[string]*schema.Schema{
},
        </code></pre>
      </div>
    </div>
  </div>
{{end}}