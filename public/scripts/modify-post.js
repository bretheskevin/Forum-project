export function modifyBtnEvent() {
    for (let btn of document.body.getElementsByClassName("modify-btn")) {
        btn.addEventListener("click", () => {
            const id = btn.parentElement.parentElement.previousSibling.previousSibling.textContent;
            window.location.href = "/edit-topic?id=" + id;
        })
    }
}
