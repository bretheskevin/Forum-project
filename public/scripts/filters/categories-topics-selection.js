import {generatePostsByCategory} from "../generate-posts/generate-posts-by-category.js";
import {changeTopics} from "../generate-posts/change-topics.js";

const categories = document.body.getElementsByClassName("categories");
const topics = document.body.getElementsByClassName("topics");

for (let category of categories) {
    category.addEventListener("click", async () => {
        changeTopics(category.getAttribute("id"), topics)
        document.getElementById("category").textContent = category.getAttribute("id");
        for (let topic of topics) {
            if (topic.classList.contains("shadow-4")) {
                document.getElementById("topic").textContent = topic.getAttribute("id");
                break;
            }
        }


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

        try {
            await generatePostsByCategory(
            document.getElementById("posts-container"),
            document.getElementsByClassName("active-primary-filter")[0].parentElement.getAttribute("id")
            )
        } catch (e) {
            // do nothing
        }

    })
}

for (let topic of topics) {
    topic.addEventListener("click", async () => {
        document.getElementById("topic").textContent = topic.getAttribute("id");

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

        try {
            await generatePostsByCategory(
                document.getElementById("posts-container"),
                document.getElementsByClassName("active-primary-filter")[0].parentElement.getAttribute("id")
            )
        } catch (e) {
            // do nothing
        }
    })
}