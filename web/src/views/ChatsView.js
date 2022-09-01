import AbstractView from "./AbstractView.js";

export default class extends AbstractView{
    constructor() {
        super();
        this.setTitle("Chats");
    }
    async getHtml(){
        //chats
        return `
        `;
    }

}