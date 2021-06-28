const startNewTopic = document.getElementById("start-new-topic");
const elementsToHideIfNotLogged = document.body.getElementsByClassName("hide-if-not-logged");

if (document.cookie === "") {
    startNewTopic.href = "/login";
    for (let element of elementsToHideIfNotLogged) {
        element.classList.add("hide")
    }
}