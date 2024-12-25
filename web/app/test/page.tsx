import Image from "next/image";

export default function Test() {
    return (
        <>
            <div>
                YOU COULD BE RIGHT YO ARE SO RIGHT
            </div>
            <Image
                src={"/jjk-godlike.jpeg"}
                alt="jjk"
                width={400}
                height={400}
            />
            <img src="/jjk-godlike.jpeg" alt="jjk" />

            <img src="/file.svg" height={20} alt="sth" />
        </>
    )
}
