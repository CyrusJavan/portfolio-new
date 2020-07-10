// Terraform Schema Tool

$(document).ready(function () {
    refreshTF();
    $("button#new-attribute-btn").click(newAttributeClickHandler);
    $("select#type").change(typeChangeHandler);
    $("#enable-default").change(defaultCheckboxHandler)
});

let attributes = [];

function newAttributeClickHandler() {
    let name = $("input#name").val();
    if (name === "") {
        return;
    }

    for (let i = 0; i < attributes.length; i++) {
        if (attributes[i].name === name) {
            alertError("Attribute with name: " + name + " already exists.");
            return;
        }
    }

    let attrType = $("select#type").val();
    let newAttr = {
        name: name,
        type: attrType,
        constraint: $("select#constraint").val(),
        forceNew: $("input#force-new").prop("checked"),
        defaultIsString: false
    };

    let enableDefault = $("input#enable-default").prop("checked");
    if (enableDefault) {
        let defaultVal;
        if (attrType === "Bool") {
            defaultVal = $("#default-value-bool").val();
        } else if (attrType === "Int" || attrType === "Float") {
            defaultVal = $("#default-value-number").val();
        } else if (attrType === "String") {
            defaultVal = $("#default-value-string").val();
            newAttr.defaultIsString = true;
        }
        newAttr.defaultVal = defaultVal;
    }

    attributes.push(newAttr);
    refreshTF();
}

function defaultCheckboxHandler() {
    let enableDefault = $("input#enable-default").prop("checked");

    if (enableDefault) {
        $("#default-form").show();
    } else {
        $("#default-form").hide();
    }
}

function typeChangeHandler() {
    let newType = $("select#type").val();
    $(".default-input").hide();

    if (newType === "Bool") {
        $("#default-value-bool").show();
    } else if (newType === "Int" || newType === "Float") {
        let numInput = $("#default-value-number");
        numInput.show();
    } else if (newType === "String") {
        $("#default-value-string").show();
    }
}

// refreshTF will render the UI with the current attributes.
function refreshTF() {
    $("#alerts").empty();
    renderAttrDisplay(attributes);
    renderCodeBox(attributes);
}

const attrTemplate = Handlebars.compile(`
<li id="listgroup-{{{name}}}" class="list-group-item">
    <p>Name: <span class="text-primary">{{{name}}}</span></p>
    <p>Type: <span class="text-primary">{{{type}}}</span></p>
    <p>Constraint: <span class="text-primary">{{{constraint}}}</span></p>
    {{#if defaultVal}}
    <p>Default: <span class="text-primary">{{{defaultVal}}}</span></p>
    {{/if}}
    <p>ForceNew: <span class="text-primary">{{{forceNew}}}</span></p>
    <p><a style="display: none;" href="#" id="trash-{{{name}}}"><i class="fa fa-trash" aria-hidden="true"></i></a></p>
</li>
`);

// renderAttrDisplay will render the Attribute Display Box.
function renderAttrDisplay() {
    let ab = $('#attribute-box');
    ab.empty();
    for (let i = 0; i < attributes.length; i++) {
        let a = attributes[i];

        let el = attrTemplate(a);
        ab.append(el);

        $(`#listgroup-${a.name}`).hover(
            function () {
                $(`#trash-${a.name}`).show();
            },
            function () {
                $(`#trash-${a.name}`).hide();
            }
        );

        $(`#trash-${a.name}`).click(function () {
            attributes = attributes.filter(function (ele) {
                return ele.name !== a.name;
            });
            refreshTF();
        });
    }
}

const codeTemplate = Handlebars.compile(`
Schema: map[string]*schema.Schema{
    {{#each attributes}}
    "{{{this.name}}}": {
        Type: schema.Type{{{this.type}}},
        {{{this.constraint}}}: true,
    {{#if this.defaultVal}}
        Default: {{#if this.defaultIsString}}"{{{this.defaultVal}}}"{{else}}{{{this.defaultVal}}}{{/if}},
    {{/if}}    
    {{#if this.forceNew}}
        ForceNew: true,
    {{/if}}
    },
    {{/each}}
},
`);

// renderCodeBox renders the schema code template.
function renderCodeBox() {
    let newCode = codeTemplate({attributes: attributes});
    $("#tf-code").text(newCode);
}

// alertError is a utility function to show an alert message to the user.
function alertError(msg) {
    let err = `
<div class="alert alert-danger alert-dismissible fade show" role="alert">
  ${msg}
  <button type="button" class="close" data-dismiss="alert" aria-label="Close">
    <span aria-hidden="true">&times;</span>
  </button>
</div>
`;
    $("#alerts").append(err);
}