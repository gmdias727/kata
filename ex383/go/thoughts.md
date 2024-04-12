dado duas strings:
ransomNote e magazine
    retorna TRUE se:
    ransomNote pode ser construído usando as letras de magazine
    do contrário retorna FALSE

portanto:
ransomNote = "aa"
magazine = "aab"
retorna TRUE porque "aab" (de magazine) pode ser construído utilizando 
das letras "aa" de ransomNote

Raciocínio:
1 - Dado duas strings
2 - Validar se:
   2.1 - O valor da primeira string existe dentro da segunda string
3 - Se existir retornar verdadeiro
4 - Caso contrário retornar falso

Exemplo:
primeira = "na"
segunda = "banana"
TRUE pois a string "na" aparece duas vezes em "banana"