// Log the date, showing JS code working from within Go
function logDate() {
  let d = date()
  console.log(
    "date type:", typeof d, "\n",
    "date is an instance of JS's Date:", d instanceof Date, "\n",
    "date value:", d.toString(), "\n"
    )
}

// Use Vue with wasm
function createApp() {
  new Vue({
    el: "#app",
    data: {
      value1: '',
      value2: '',
      result: ''
    },
    methods: {
      add: function () {
        this.result = add(this.value1, this.value2)
      },
      subtract: function () {
        this.result = subtract(this.value1, this.value2)
      }
    }
  })
}

async function logGetRequest() {
  try {
      const response = await getRequest('https://api.taylor.rest/')
      const message = await response.json()
      console.log(message)
  } catch (err) {
      console.error('Caught exception', err)
  }
}

async function main() {
    const go = new Go();
    const response = await fetch("lib.wasm");
    const buffer = await response.arrayBuffer();
    const result = await WebAssembly.instantiate(buffer, go.importObject);
    go.run(result.instance)

    logDate()
    createApp()
    await logGetRequest()
}

(async function() {
  await main();
})();