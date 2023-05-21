import Link from "next/link";
import { Fragment } from "react";

type threadProps = {
    id: number;
    title: string;
    text: string;
    userId: number;
    channelId: number;
}

const Thread = (thread: threadProps) => {
    return (
        <Fragment key={thread.id}>
        <section className="text-gray-600 body-font">
        <div className="container px-5 py-24 mx-auto flex flex-wrap">
            <h2 className="sm:text-3xl text-2xl text-gray-900 font-medium title-font mb-2 md:w-2/5">{thread.title}</h2>
            <div className="md:w-3/5 md:pl-6">
            <p className="leading-relaxed text-base text-ellipsis overflow-hidden">{thread.text}</p>
            <div className="flex md:mt-4 mt-6">
                <Link href={`/threads/${thread.id}`}>
                    <button type="button" className="inline-flex text-white bg-indigo-500 border-0 py-1 px-4 focus:outline-none hover:bg-indigo-600 rounded">Открыть</button>
                </Link>
                <a className="text-indigo-500 inline-flex items-center ml-4">
                <svg fill="none" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" className="w-4 h-4 ml-2" viewBox="0 0 24 24">
                    <path d="M5 12h14M12 5l7 7-7 7"></path>
                </svg>
                </a>
            </div>
            </div>
        </div>
        </section>
        </Fragment>
    );
}

export default Thread;