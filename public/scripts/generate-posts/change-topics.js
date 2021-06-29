export function changeTopics(category, topics) {
    let topicsList = [];
    let topicsAttr = [];
    switch (category) {
        case "technology":
            topicsList = ["Computer", "Crypto currency", "Cyber-security", "Web Development", "Python", "Network", "Gaming"];
            topicsAttr = ["computer", "crypto_currency", "cyber_security", "web_development", "python", "network", "gaming"];
            for (let i = 0; i < topics.length; i++) {
                topics[i].setAttribute("id", topicsAttr[i])
                topics[i].lastChild.textContent = topicsList[i]
            }
            break;
        case "social":
            topicsList = ["Climate Change", "Overpopulation", "Civil Rights", "Gender Inequality", "Health Care", "Migration", "Fast foods"];
            topicsAttr = ["climate_change", "overpopulation", "civil_rights", "gender_inequality", "health_care", "migration", "fast_foods"];
            for (let i = 0; i < topics.length; i++) {
                topics[i].setAttribute("id", topicsAttr[i])
                topics[i].lastChild.textContent = topicsList[i]
            }
            break;
        case "reader":
            topicsList = ["Biographies", "Religion", "Fitness", "Cookbooks", "Business", "Education", "Horror"];
            topicsAttr = ["biographies", "religion", "fitness", "cookbooks", "business", "education", "horror"];
            for (let i = 0; i < topics.length; i++) {
                topics[i].setAttribute("id", topicsAttr[i])
                topics[i].lastChild.textContent = topicsList[i]
            }
            break;
        case "sports":
            topicsList = ["Tennis", "Football", "Rugby", "Swimming", "Volley", "Handball", "Basket"];
            topicsAttr = ["tennis", "football", "rugby", "swimming", "volley", "handball", "basket"];
            for (let i = 0; i < topics.length; i++) {
                topics[i].setAttribute("id", topicsAttr[i])
                topics[i].lastChild.textContent = topicsList[i]
            }
            break;
        case "political":
            topicsList = ["Vaccines", "Abortion", "Animal Rights", "Global Climate Change", "Cancel Culture", "Marriage Equality", "Black Lives Matter"];
            topicsAttr = ["vaccines", "abortion", "animal_rights", "global_climate_change", "cancel_culture", "marriage_equality", "black_lives_matter"];
            for (let i = 0; i < topics.length; i++) {
                topics[i].setAttribute("id", topicsAttr[i])
                topics[i].lastChild.textContent = topicsList[i]
            }
            break;
    }
}