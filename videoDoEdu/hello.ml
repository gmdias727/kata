let x = 1 + 2

(* let add = fun (a: int) (b: int) -> a + b
let add = fun (a: int) -> fun (b: int) -> a + b
let add = fun ((a: int), (b: int)) -> a + b *)

let add (a: int) (b: int) = a + b

let three = add 1 2

(* let message_size n =
  let element_size = 4 in
  let separator_size = 1 in
  let constant = element_size + separator_size in
  constant * n *)

let message_size n =
  let constant = 
    let element_size = 4 in
    let separator_size = 1 in
    element_size + separator_size
  in
  Format.printf "constant: %d\n" constant;
  constant * n


type status = Valid | Banned of {reason : string} | Expired
type user = { name: string; status : status; }
let print_user user =
  let {name; status} = user in
  match status with
  | Valid -> Format.printf "%s is a good user" name
  | Expired -> Format.printf "%s is a expired user" name
  | Banned { reason } -> Format.printf "%s is a bad user, because %s" name reason

let () = print_user { name = "Grande"; status = Banned { reason = "amou demais \n"} }
let () = Format.printf "aaa\n"

let arquivo = "demo.txt"
let message = "Hello World!"

let () = 
  let oc = open_out arquivo in
  Printf.fprintf oc "%s\n" message;
  close_out oc;

  let ic = open_in arquivo in
  try
    let line = input_line ic in
    print_endline line;
    flush stdout;
    close_in ic
  with e ->
    close_in_noerr ic;
    raise e