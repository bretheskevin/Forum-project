const postId = parseInt(window.location.href.split("?")[1].split("=")[1]);
const titleContent = document.getElementById("title");
const postContent = document.getElementById("content");

function activeCategory(element) {
    const classes = "shadow-4 p-3 rounded-2 white".split(" ");
    for (let category of document.body.getElementsByClassName("categories")) {
        for (let className of classes) {
            category.classList.remove(className);
        }
    }
    for (let className of classes) {
        element.classList.add(className);
    }
}

function activeTopics(element) {
    const classes = "shadow-4 p-3 rounded-2 white".split(" ");
    for (let topic of document.body.getElementsByClassName("topics")) {
        for (let className of classes) {
            topic.classList.remove(className);
        }
    }
    for (let className of classes) {
        element.classList.add(className);
    }
}

async function fetchData() {
    const response = await fetch("/post/" + postId);
    return await response.json();
}

async function main() {
    const data = await fetchData();
    const category = data["Category"].split("-")[0];
    const topic = data["Category"].split("-")[1];

    titleContent.value = data["Title"];
    postContent.value = data["Content"];
    activeCategory(document.getElementById(category))
    document.getElementById("category").textContent = category;
    activeTopics(document.getElementById(topic));
    document.getElementById("topic").textContent = topic;
}

main()
