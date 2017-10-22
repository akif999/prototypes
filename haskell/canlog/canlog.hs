import System.IO
import Data.Text.IO as DTI

get_and_print :: Handle -> IO ()
get_and_print fp = do
    eof <- hIsEOF fp
    if eof then return ()
    else do
        str <- DTI.hGetLine fp
        print str
        get_and_print fp

main :: IO ()
main = do
    fp <- openFile "srec_sample_s1_01.txt" ReadMode
    get_and_print fp
