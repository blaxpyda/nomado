import { API } from "./services/API.js";

window.app = {
    search: (event) => {
        event.preventDefaukt();
        const q = document.querySelector("input[type='search']").value;
        // TODO
    },
    api: API
}