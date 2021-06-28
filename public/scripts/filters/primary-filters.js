import {generatePostsByCategory} from "../generate-posts/generate-posts-by-category.js";

const container = document.getElementById("posts-container");

const filters = document.body.getElementsByClassName("filters");
for (let filter of filters) {
    filter.addEventListener("click", () => {
        for (let filt of filters) {
            filt.firstElementChild.classList.remove("bg-primary");
            filt.firstElementChild.classList.remove("txt-white");
            filt.firstElementChild.classList.remove("txt-center");
            filt.firstElementChild.classList.remove("mt-5");
            filt.firstElementChild.classList.remove("mb-4");
            filt.firstElementChild.classList.remove("pt-3");
            filt.firstElementChild.classList.remove("pb-3");
            filt.firstElementChild.classList.remove("rounded-2");
            filt.firstElementChild.classList.remove("font-w700");
            filt.firstElementChild.classList.remove("active-primary-filter");
        }
        filter.firstElementChild.classList.add("bg-primary");
        filter.firstElementChild.classList.add("txt-white");
        filter.firstElementChild.classList.add("txt-center");
        filter.firstElementChild.classList.add("mt-5");
        filter.firstElementChild.classList.add("mb-4");
        filter.firstElementChild.classList.add("pt-3");
        filter.firstElementChild.classList.add("pb-3");
        filter.firstElementChild.classList.add("rounded-2");
        filter.firstElementChild.classList.add("font-w700");
        filter.firstElementChild.classList.add("active-primary-filter");
    })
}

document.getElementById("latest").addEventListener("click", async () => {
    await generatePostsByCategory(container, "latest");
})

document.getElementById("older").addEventListener("click", async () => {
    await generatePostsByCategory(container, "older");
})
