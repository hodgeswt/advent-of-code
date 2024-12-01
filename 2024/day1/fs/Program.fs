open System.IO

module Day1 =
    let Run path =
        File.ReadLines path
        |> Seq.map(fun x -> x.Split("   "))
        |> Seq.map(fun x -> x |> Array.map int)
        |> Seq.map(fun x -> (x.[0], x.[1]))
        |> Seq.toArray
        |> Array.unzip
        |> fun (x, y) -> (x |> Array.sort, y |> Array.sort)
        |> fun (x, y) -> Array.zip x y
        |> Array.map(fun (x, y) -> abs(x - y))
        |> Array.sum
        |> printfn "%d"

[<EntryPoint>]
let main argv =
    if argv.Length = 0 then
        // Run actual data
        Day1.Run "../day1.input"
        0
    else
        if argv.[0] <> "-t" then
            printfn "Only accepted argument is -t"
            1
        else
            // Run tests
            Day1.Run  "../day1.test"
            0



