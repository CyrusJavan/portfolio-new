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