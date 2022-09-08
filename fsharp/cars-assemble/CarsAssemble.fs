module CarsAssemble

let successRate (speed: int): float =
    match speed with
    | 1 | 2 | 3 | 4 -> 1.0
    | 5 | 6 | 7 | 8 -> 0.9
    | 9  -> 0.8
    | 10 -> 0.77
    | _  -> 0

let carsPerHourBaseline = 221

let productionRatePerHour (speed: int): float =
    float speed * float carsPerHourBaseline * successRate speed

let workingItemsPerMinute (speed: int): int =
    let productionRate = productionRatePerHour speed
    int productionRate / 60
