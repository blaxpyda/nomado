export const API = {
    baseURL: "/api/v1",
    getTopProperty: async () => {
       return await API.fetch("properties/top");
    },
    getRandomProperty: async () => {
       return await API.fetch("properties/random");
    },
    getPropertyById : async (id) => {
       return await API.fetch(`properties/${id}`);
    },
    searchProperty: async (q) => {
       return await API.fetch(`properties/search?q=${encodeURIComponent(q)}`);
    },
    fetch: async (serviceName, args) => {
        try {
            const response = await fetch(API.baseURL + serviceName);
            const result = await response.json();
            return result;

        } catch (e) {
            console.error(e);
        }
        

    }
}