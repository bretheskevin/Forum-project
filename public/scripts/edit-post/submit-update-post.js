import {Api} from "../api/api.js";

const parseJwt = (token) => {
    try {
        return JSON.parse(atob(token.split('.')[1]));
    } catch (e) {
        return null;
    }
}

document.getElementById("form-submit-btn").addEventListener("click", () => {
    const title = document.getElementById("title").value;
    const content = document.getElementById("content").value;
    const category = document.getElementById("category").textContent;
    const topic = document.getElementById("topic").textContent;
    const loggedUserId = parseInt(parseJwt(document.cookie.split("=")[1])["user_id"]);
    const postId = parseInt(document.getElementById("post-id").textContent);

    new Api().updatePost(postId, title, content, loggedUserId, category, topic);
});