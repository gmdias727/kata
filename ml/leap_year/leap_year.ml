let () =
  try
    let year = read_int () in
    let leap = (year mod 4 = 0 && year mod 100 <> 0) || year mod 400 = 0 in
    if leap then
      Printf.printf "%d is a leap year\n" year
    else
      Printf.printf "%d is not a leap year\n" year
with
| Failure _ -> Printf.printf "Please enter a valid integer"