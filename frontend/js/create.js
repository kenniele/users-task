let createButton = document.getElementById("createForm").querySelector("button")
if (createButton) {
    createButton.onclick = function (e) {
        e.preventDefault()
        let inputs = document.getElementById("createForm").querySelectorAll("input")
        let block = document.getElementById("createForm").querySelector("#block")
        let data = {}
        console.log(inputs)

        inputs.forEach(input => {
            data[input.name] = input.value
        })

        console.log(data)
        console.log(data.passportNumber)

        let xhr = new XMLHttpRequest()
        xhr.open("POST", `/create/${data.passportNumber}`)
        xhr.setRequestHeader("Content-Type", "application/json")

        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response)
            if (response && "Error" in response) {
                if (response.Error == null) {
                    block.innerText = "Аккаунт успешно создан"
                    setTimeout(() => {
                        block.innerText = ""
                    }, 1000)
                } else {
                    block.innerText = `Возникла ошибка ${response.Error}`
                    setTimeout(() => {
                        block.innerText = ""
                    }, 1000)
                }
            } else {
                block.innerText = "Некорректные данные"
                setTimeout(() => {
                    block.innerText = ""
                }, 1000)
            }
        }
        xhr.send(JSON.stringify(data))
    }
}