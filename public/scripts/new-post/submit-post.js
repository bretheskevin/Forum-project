import {Api} from "../api/api.js";

document.getElementById("form-submit-btn").addEventListener("click", () => {
    const title = document.getElementById("title").value;
    const content = document.getElementById("content").value;
    const category = document.getElementById("category").textContent;
    const topic = document.getElementById("topic").textContent;

    new Api().uploadNewPost(title, content, category, topic)
});