const categories = document.body.getElementsByClassName("categories");

for (let category of categories) {
    category.addEventListener("click", () => {
        const id = category.getAttribute("id");

        document.getElementById("category").textContent = id;

        for (let categ of categories) {
            categ.classList.remove("white");
            categ.classList.remove("shadow-4");
            categ.classList.remove("p-3");
            categ.classList.remove("rounded-2");
        }
        category.classList.add("white");
        category.classList.add("shadow-4");
        category.classList.add("p-3");
        category.classList.add("rounded-2");
    })
}


const topics = document.body.getElementsByClassName("topics");

for (let topic of topics) {
    topic.addEventListener("click", () => {
        const id = topic.getAttribute("id");
        document.getElementById("topic").textContent = id;

        for (let top of topics) {
            top.classList.remove("white");
            top.classList.remove("shadow-4");
            top.classList.remove("p-3");
            top.classList.remove("rounded-2");
        }
        topic.classList.add("white");
        topic.classList.add("shadow-4");
        topic.classList.add("p-3");
        topic.classList.add("rounded-2");
    })
}