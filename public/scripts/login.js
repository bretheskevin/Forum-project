if (document.cookie) {
    window.location.href = "/feed";
}

let validEmail = false;
let validPassword = false;

function isValid() {
    const btn = document.getElementById("form-submit-login");
    if (validEmail && validPassword) {
        btn.classList.remove("disabled");
    } else {
        btn.classList.add("disabled");
        $(document).keypress(
            function(event) {
                if (event.which === '13') {
                    event.preventDefault();
                }
            });
    }
}


// email

function validateEmail(email) {
    const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}

const email = document.getElementById("email")
email.addEventListener("input", () => {
    const invalidEmail = document.getElementById("invalid-email");

    if (validateEmail(email.value)) {
        invalidEmail.classList.add("hide");
        validEmail = true;
    } else {
        invalidEmail.classList.remove("hide");
        validEmail = false;
    }

    isValid();
})

// password
const passwordShort = document.getElementById("password-short");
const passwordLong = document.getElementById("password-long");

const password = document.getElementById("password");
password.addEventListener("input", () => {
    if (password.value.length < 8) {
        passwordShort.classList.remove("hide");
    } else {
        passwordShort.classList.add("hide");
    }
    if (password.value.length > 30) {
        passwordLong.classList.remove("hide");
    } else {
        passwordLong.classList.add("hide");
    }

    validPassword = password.value.length >= 8 && password.value.length <= 30;

    isValid();
})

const submitBtn = document.getElementById("form-submit-login");
submitBtn.addEventListener("click", () => {
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    const request = new XMLHttpRequest();
    request.onreadystatechange = function() {
        if (this.readyState === 4 && this.status === 200) {
            if (this.responseText.includes("Wrong")) {
                document.getElementById("error").classList.remove("hide")
            } else {
                window.location.href = "/feed"
            }
        }
    };

    request.open("POST", "/login");
    request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    request.responseType = "text";
    request.send(JSON.stringify({
        "email": email,
        "password": password
    }));


    // change the color of the btn
    submitBtn.classList.remove("black");
    submitBtn.classList.add("grey");
    setTimeout((() => {
        submitBtn.classList.remove("grey")
        submitBtn.classList.add("black")
    }),
        100)
})
