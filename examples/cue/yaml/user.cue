import "time"

#User: {
    id: int
    name: string
    created: time.Format("2006-01-02")
}

john: #User
john: {
    id: 1
    name: "John"
    created: "2023-02-09"
}
