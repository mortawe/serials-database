export const path = "/components/Form/"

export class FormComponent {
    #parent
    #data
    #tmplName
    #config

    constructor({
                    tmplName = '',
                    parent = document.body,
                    config = [],
                    data = []
                } = {}) {
        this.#parent = parent;
        this.#data = data;
        this.#tmplName = tmplName;
        this.#config = config;
    }

    render() {
        switch (this.#tmplName) {
            case '':
                console.error('unexpected table type')
                break;
            default:
                this.#renderTmpl(path + this.#tmplName);
                break;
        }
    }

    #renderTmpl(tmplPath) {
        fetch(tmplPath)
            .then(response => response.text())
            .then((template) => {
                this.#parent.innerHTML = Mustache.render(template, {config: this.#config});
            });
    }
}
