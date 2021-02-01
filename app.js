// Use Vue with wasm
function createApp() {
  new Vue({
    el: "#app",
    data: {
      value1: '',
      value2: '',
      result: '',
      quote: ''
    },
    methods: {
      add: function () {
        this.result = add(this.value1, this.value2)
      },
      subtract: function () {
        this.result = subtract(this.value1, this.value2)
      },
      getQuote: function () {
        let boundThis = this;

        (async function () {
          try {
            const response = await getTaylorSwiftQuote('https://api.taylor.rest/')
            const resJson = await response.json()
            boundThis.quote = resJson['quote']
          } catch (err) {
            console.log(err)
            return err
          }
        })()
      }
    }
  })
}

async function main() {
  const go = new Go();
  const response = await fetch("lib.wasm");
  const buffer = await response.arrayBuffer();
  const result = await WebAssembly.instantiate(buffer, go.importObject);
  go.run(result.instance)

  createApp()
}

(async function () {
  await main();
})();