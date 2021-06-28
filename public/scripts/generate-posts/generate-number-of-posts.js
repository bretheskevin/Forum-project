export function generateNumberOfPosts(postsList){
    switch (postsList.length) {
        case 0:
            document.getElementById("nb-of-posts").textContent = postsList.length + " post"
            break;
        case 1:
            document.getElementById("nb-of-posts").textContent = postsList.length + " post"
            break;
        default:
            document.getElementById("nb-of-posts").textContent = postsList.length + " posts"
            break;
    }
}