const chevrons = document.body.getElementsByClassName("chevrons");
const triggers = document.body.getElementsByClassName("collapsible-trigger");

for (const trigger of triggers) {
    trigger.addEventListener("click", () => {
        for (const chevron of chevrons) {
            if (trigger.lastElementChild === chevron) {
                if (chevron.classList.contains("mdi-chevron-down")) {
                    chevron.classList.remove("mdi-chevron-down");
                    chevron.classList.add("mdi-chevron-up");
                } else {
                    chevron.classList.add("mdi-chevron-down");
                    chevron.classList.remove("mdi-chevron-up");
                }
            }
        }
    })
}

