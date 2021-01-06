const squareWasm = (input) => {
  output.value = square(input)
}

const input = document.getElementById('input')
const output = document.getElementById('output')

const onFormat = () => {
  const string = formatJSON(input.value)
  output.value = string
}

const onClear = () => {
  input.value = ''
}

document.addEventListener('click', (event) => {
  const { id } = event.target
  if (id === 'format') {
    onFormat()
  } else if (id === 'clear') {
    onClear()
  }
})
