import System.IO
import System.Environment

main :: IO()
main = do
    args <- getArgs
    handle <- openFile (args !! 0) ReadMode
    text <- hGetContents handle
    putStr text
    hClose handle
