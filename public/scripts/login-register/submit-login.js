import {Api} from "../api/api.js";

const submitBtn = document.getElementById("form-submit-login");
submitBtn.addEventListener("click", () => {
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    new Api().submitLogin(email, password);
})
