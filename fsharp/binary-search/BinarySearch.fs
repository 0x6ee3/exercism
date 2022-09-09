module BinarySearch

let find input value =
    input |> Seq.tryFindIndex (fun v -> v = value)
