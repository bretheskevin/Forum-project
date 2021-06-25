export function deleteBtnEvent() {
    for (let btn of document.body.getElementsByClassName("modify-btn")) {
        btn.addEventListener("click", () => {
            const id = btn.parentElement.parentElement.previousSibling.previousSibling.textContent;

            const xhr = new XMLHttpRequest();
            xhr.onreadystatechange = function() {
                if (this.readyState === 4 && this.status === 200) {
                    window.location.reload();
                }
            };

            xhr.open("PATCH", "/post/" + id);
            xhr.send();
        })
    }
}
