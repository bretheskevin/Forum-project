export class Api {
    deletePostById(id) {
        const xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function() {
            if (this.readyState === 4 && this.status === 200) {
                window.location.reload();
            }
        };

        xhr.open("DELETE", "/post/" + id);
        xhr.send();
    }

    async getPostsByCategory(category) {
        // return a list containing every posts

        const res = await fetch("/posts/" + category)
        return await res.json();
    }

    async getPostsByUser(userId) {
        const res = await fetch("/posts/user/" + userId)
        return await res.json();
    }

    uploadNewPost(title, content, category, topic) {
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
    }

    submitRegister(email, username, password) {
        const request = new XMLHttpRequest();
        request.onreadystatechange = function() {
            if (this.readyState === 4 && this.status === 200) {
                const response = this.responseText;
                const error = document.getElementById("error");
                const errorMessage = document.getElementById("error-content");
                if (response.includes("username")) {
                    error.classList.remove("hide");
                    errorMessage.textContent = "The username is already taken !";
                } else if (response.includes("email")) {
                    error.classList.remove("hide");
                    errorMessage.textContent = "The email is already taken !";
                } else {
                    window.location.href = "/feed"
                }
            }
        };

        request.open("POST", "/register");
        request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        request.responseType = "text";
        request.send(JSON.stringify({
            "email": email,
            "username": username,
            "password": password
        }));
    }


    submitLogin(email, password) {
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
    }

    updatePost(postId, title, content, loggedUserId, category, topic) {
        const xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function() {
            if (this.readyState === 4 && this.status === 200) {
                window.location.href = "/dashboard/posts"
            }
        };

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
    }
}