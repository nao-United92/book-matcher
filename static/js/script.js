// パスワード表示・非表示切り替え
function togglePasswordVisibility() {
  var passwordField = document.getElementById("password");
  var passwordToggle = document.querySelector(".display-pwd-on");

  if (passwordField.type === "password") {
      passwordField.type = "text";
      passwordToggle.textContent = "visibility";
  } else {
      passwordField.type = "password";
      passwordToggle.textContent = "visibility_off";
  }
}
