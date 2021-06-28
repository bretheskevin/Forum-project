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

    const xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState === 4 && this.status === 200) {
            window.location.href = "/dashboard/posts"
        }
    };

    console.log(postId)
    xhr.open("PATCH", "/post");
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.responseType = "text";
    xhr.send(JSON.stringify({
        "id": postId,
        "title": title,
        "content": content,
        "publisherId": loggedUserId,
        "category": category+"-"+topic
    }));

});