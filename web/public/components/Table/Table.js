export const path = "/components/Table/"

export class TableComponent {
    #parent
    #data
    #tmplName

    constructor({
                    tmplName = '',
                    parent = document.body,
                    data = []

                } = {}) {
        this.#parent = parent;
        this.#data = data;
        this.#tmplName = tmplName;
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

    #renderTmpl(tmpPath) {
        fetch(tmpPath)
            .then(response => response.text())
            .then((template) => {
                this.#parent.innerHTML = Mustache.render(template, {data: this.#data});
            });
    }
}