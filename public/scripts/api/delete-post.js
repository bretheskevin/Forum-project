import {Api} from "./api.js";
const api = new Api()

export function deleteBtnEvent() {
    for (let btn of document.body.getElementsByClassName("delete-btn")) {
        btn.addEventListener("click", () => {
            const id = btn.parentElement.parentElement.previousSibling.previousSibling.textContent;
            api.deletePostById(id);
        })
    }
}
