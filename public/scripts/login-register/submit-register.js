import {Api} from "../api/api.js";

const submitBtn = document.getElementById("form-submit-register");
submitBtn.addEventListener("click", () => {
    const email = document.getElementById("email").value;
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    new Api().submitRegister(email, username, password);
    // change the color of the btn
    submitBtn.classList.remove("black");
    submitBtn.classList.add("grey");
    setTimeout((() => {
            submitBtn.classList.remove("grey")
            submitBtn.classList.add("black")
        }),
        100)
})