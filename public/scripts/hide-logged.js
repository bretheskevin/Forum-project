const elementsToHideIfLogged = document.body.getElementsByClassName("hide-if-logged");

if (document.cookie) {
    for (let element of elementsToHideIfLogged) {
        element.classList.add("hide")
    }
}