// Terraform Schema Tool

$(document).ready(function () {
    initDataStructures();
    initHandlers();
    refreshTF();
});

let rootResource;

let aggregateTypes = ["Set", "List", "Map"];
let numericTypes = ["Int", "Float"];

function initDataStructures() {
    rootResource = new ResourceNode({name: "_root"});
}

function getTopLevelAttributeValues() {
    let values = [];
    for (let i = 0; i < rootResource.children.length; i++) {
        values.push(rootResource.children[i].value);
    }
    return values;
}

function initHandlers() {
    $("button#new-attribute-btn").click(newAttributeClickHandler);
    $("select#type").change(typeChangeHandler);
    $("select#elem-type").change(elemTypeChangeHandler);
    $("#enable-default").change(defaultCheckboxHandler);
}

function newAttributeClickHandler() {
    let name = $("input#name").val();
    if (name === "") {
        return;
    }

    let attributes = getTopLevelAttributeValues();
    for (let i = 0; i < attributes.length; i++) {
        if (attributes[i].name === name) {
            alertError("Attribute with name: " + name + " already exists.");
            return;
        }
    }

    let attrType = $("select#type").val();
    let elemType = $("select#elem-type").val();
    let newAttr = {
        name: name,
        type: attrType,
        constraint: $("select#constraint").val(),
        forceNew: $("input#force-new").prop("checked"),
        defaultIsString: false,
        elemType: ""
    };

    let enableDefault = $("input#enable-default").prop("checked");
    if (enableDefault) {
        let defaultVal;
        if (attrType === "Bool") {
            defaultVal = $("#default-value-bool").val();
        } else if (numericTypes.includes(attrType)) {
            defaultVal = $("#default-value-number").val();
        } else if (attrType === "String") {
            defaultVal = $("#default-value-string").val();
            newAttr.defaultIsString = true;
        } else if (aggregateTypes.includes(attrType)) {
            if (elemType === "Bool") {
                defaultVal = $("#default-value-bool").val();
            } else if (numericTypes.includes(elemType)) {
                defaultVal = $("#default-value-number").val();
            } else if (elemType === "String") {
                defaultVal = $("#default-value-string").val();
                newAttr.defaultIsString = true;
            }
        }
        newAttr.defaultVal = defaultVal;
    }

    if (aggregateTypes.includes(attrType)) {
        newAttr.elemType = elemType;
    }

    let newNode = new ResourceNode(newAttr);
    rootResource.children.push(newNode);
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
    $("#elem-type-form-group").hide();

    if (newType === "Bool") {
        $("#default-value-bool").show();
    } else if (numericTypes.includes(newType)) {
        let numInput = $("#default-value-number");
        numInput.show();
    } else if (newType === "String") {
        $("#default-value-string").show();
    } else if (aggregateTypes.includes(newType)) {
        $("#elem-type-form-group").show();
        elemTypeChangeHandler();
    }
}

function elemTypeChangeHandler() {
    $(".default-input").hide();

    let elemType = $("select#elem-type").val();
    if (elemType === "Bool") {
        $("#default-value-bool").show();
    } else if (numericTypes.includes(elemType)) {
        let numInput = $("#default-value-number");
        numInput.show();
    } else if (elemType === "String") {
        $("#default-value-string").show();
    }
}

// refreshTF will render the UI with the current attributes.
function refreshTF() {
    $("#alerts").empty();
    renderAttrDisplay();
    renderCodeBox();
}

const attrTemplate = Handlebars.compile(`
<li id="listgroup-{{{name}}}" class="list-group-item">
    <div class="d-flex">
      <div class="col-count-2">
            <p>Name: <span class="text-primary">{{{name}}}</span></p>
            <p>Type: <span class="text-primary">{{{type}}}</span></p>
            {{#if elemType}}
              <p>Elem Type: <span class="text-primary">{{{elemType}}}</span></p>
            {{/if}}
            <p>Constraint: <span class="text-primary">{{{constraint}}}</span></p>
            {{#if defaultVal}}
              <p>Default: <span class="text-primary">{{{defaultVal}}}</span></p>
            {{/if}}
            <p>ForceNew: <span class="text-primary">{{{forceNew}}}</span></p>
        </div>
        <div class="ml-auto">
          <p><a href="#" id="trash-{{{name}}}"><i class="fa fa-trash" aria-hidden="true"></i></a></p>
          <p><a href="#" id="edit-{{{name}}}"><i class="fa fa-edit" aria-hidden="true"></i></a></p>
        </div>
    </div>
</li>
`);

// renderAttrDisplay will render the Attribute Display Box.
function renderAttrDisplay() {
    let ab = $('#attribute-box');
    ab.empty();
    let attributes = getTopLevelAttributeValues();
    for (let i = 0; i < attributes.length; i++) {
        let a = attributes[i];

        let el = attrTemplate(a);
        ab.append(el);

        $(`#trash-${a.name}`).click(function () {
            rootResource.children = rootResource.children.filter(function (ele) {
                return ele.value.name !== a.name;
            });
            refreshTF();
        });

        $(`#edit-${a.name}`).click(function () {
            console.log("Edit");
        });
    }
}

Handlebars.registerHelper('empty', function (l) {
    return l.length === 0
})

const codeTemplate = Handlebars.compile(
    `func resource() *schema.Resource {
	return &schema.Resource{
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
            {{#if this.elemType}}
                Elem: &schema.Schema{
                    Type: schema.Type{{{this.elemType}}},
                },
            {{/if}}
            },
            {{/each}}
            {{#if (empty attributes) }}
            // Happy Terraforming :)
            {{/if}}
        },
    }
}`
);

// renderCodeBox renders the schema code template.
function renderCodeBox() {
    let attributes = getTopLevelAttributeValues();
    let newCode = codeTemplate({attributes: attributes});
    newCode = gofmt(newCode);
    $("#tf-code").text(newCode);
    window.Prism.highlightAll()
}
