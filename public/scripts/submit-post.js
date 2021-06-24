document.getElementById("form-submit-btn").addEventListener("click", () => {
    const title = document.getElementById("title").value;
    const content = document.getElementById("content").value;
    const category = document.getElementById("category").textContent;
    const topic = document.getElementById("topic").textContent;

    const xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState === 4 && this.status === 200) {
            window.location.href = "/feed"
        }
    };

    xhr.open("POST", "/post");
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.responseType = "text";
    xhr.send(JSON.stringify({
        "title": title,
        "content": content,
        "category": category,
        "topic": topic,
    }));

});