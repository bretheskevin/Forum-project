const collapsibles = document.body.getElementsByClassName("collapsible");



for (const collapsible of collapsibles) {

    collapsible.addEventListener("ax.collapsible.open", () => {
        const chevron = collapsible.previousElementSibling.lastElementChild;
        if (chevron.classList.contains("mdi-chevron-down")) {
            chevron.classList.remove("mdi-chevron-down");
            chevron.classList.add("mdi-chevron-up");
        }
    });

    collapsible.addEventListener("ax.collapsible.close", () => {
        const chevron = collapsible.previousElementSibling.lastElementChild;
        if (chevron.classList.contains("mdi-chevron-up")) {
            chevron.classList.remove("mdi-chevron-up");
            chevron.classList.add("mdi-chevron-down");
        }
    });
}

for (const trigger of document.body.getElementsByClassName("collapsible-trigger")) {
    trigger.dispatchEvent(new Event("click"));
}