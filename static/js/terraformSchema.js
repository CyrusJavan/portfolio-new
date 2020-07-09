let attributes = [
    // {
    //     name: "example",
    //     type: "Bool",
    //     constraint: "Required"
    // }
];

$(document).ready(function () {
    refreshTF();
    $("button#new-attribute-btn").click(newAttributeClickHandler);
    $("select#type").change(typeChangeHandler);
});

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

    let newAttr = {
        name: name,
        type: $("select#type").val(),
        constraint: $("select#constraint").val(),
        forceNew: $("input#force-new").prop("checked")
    };
    attributes.push(newAttr);
    refreshTF();
}

function typeChangeHandler() {
    let newType = $("select#type").val();
    
}

function refreshTF() {
    $("#alerts").empty();
    fillAttrBox(attributes);
    fillCodeBox(attributes);
    console.log(attributes);
}

function fillAttrBox() {
    let ab = $('#attribute-box');
    ab.empty();
    for (let i = 0; i < attributes.length; i++) {
        let a = attributes[i];

        let el = `<li id="listgroup-${a.name}" class="list-group-item">` +
            `<p>Name: <span class="text-primary">${a.name}</span></p>` +
            `<p>Type: <span class="text-primary">${a.type}</span></p>` +
            `<p>Constraint: <span class="text-primary">${a.constraint}</span></p>` +
            `<p>ForceNew: <span class="text-primary">${a.forceNew}</span></p>` +
            `<p><a style="display: none;" href="#" id="trash-${a.name}"><i class="fa fa-trash" aria-hidden="true"></i></a></p>` +
            `</li>`;
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
        Type: schema.Type{{{this.type}}}
        {{{this.constraint}}}: true,
    {{#if this.forceNew}}
        ForceNew: true,
    {{/if}}
    },
    {{/each}}
},
`);

function fillCodeBox() {
    let newCode = codeTemplate({attributes: attributes});
    $("#tf-code").text(newCode);
}

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