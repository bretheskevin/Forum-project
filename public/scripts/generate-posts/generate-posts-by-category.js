import {likeAndDislike} from "../posts/like-dislike.js";
import {minify} from "../posts/posts-minify.js";
import {createPost} from "../posts/build-post.js";
import {generateNumberOfPosts} from "./generate-number-of-posts.js";
import {Api} from "../api/api.js";

export async function generatePostsByCategory(container, filter) {
    /*
    filters:
        - latest
        - older
     */

    container.innerHTML = "";
    let category = document.getElementById("category").textContent + "-" + document.getElementById("topic").textContent;
    const apiResponse = await new Api().getPostsByCategory(category);
    let postsList;
    switch (filter) {
        case "latest":
            postsList = apiResponse.reverse();
            break;
        case "older":
            postsList = apiResponse;
            break;
        default:
            console.error("Invalid filter !")
    }
    generateNumberOfPosts(postsList);
    for (let post of postsList) {
        await createPost(post, container);
    }
    likeAndDislike();
    minify();
}





