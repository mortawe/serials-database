-- INIT TABLES --
CREATE TABLE person
(
    person_id SERIAL PRIMARY KEY,
    name      VARCHAR(256) NOT NULL,
    birthdate DATE         NOT NULL,
    bio       TEXT         DEFAULT '',
    awards    VARCHAR(256) DEFAULT ''
);

CREATE TABLE show
(
    show_id     SERIAL PRIMARY KEY,
    title       VARCHAR(256)           NOT NULL,
    release     DATE                   NOT NULL,
    description TEXT         DEFAULT '',
    episode_num INT          DEFAULT 0 NOT NULL,
    genre       VARCHAR(256) DEFAULT ''
);

CREATE TABLE person_show
(
    person_id INT,
    FOREIGN KEY (person_id) REFERENCES person (person_id) ON DELETE CASCADE,
    show_id   INT,
    FOREIGN KEY (show_id) REFERENCES show (show_id) ON DELETE CASCADE,
    CONSTRAINT person_show_uq UNIQUE (person_id, show_id)
);

-- CREATE INDEX --

CREATE INDEX person_name_index ON person(person_id) INCLUDE (name);
CREATE INDEX person_name_index ON person(name);
CREATE INDEX show_title_index ON show(title);
CREATE INDEX person_birth_index ON person(birthdate);
CREATE INDEX show_release_index ON show(release);

-- INSERT TEST DATA --

INSERT INTO show (title, release, description, episode_num, genre)
VALUES ('Psycho-Pass', '2012-10-11',
        'Psycho-Pass is set in a futuristic Japan governed by the Sibyl System (シビュラシステム, Shibyura Shisutemu), a powerful bio-mechanical computer network which endlessly measures the biometrics of Japanese citizens'' brains and mentalities using a "cymatic scan." The resulting assessment is called a Psycho-Pass (サイコパス, Saikopasu), which includes a numeric Crime Coefficient (犯罪係数, Hanzaikeisū) index, revealing the citizen''s criminality potential, and a color-coded Hue, alerting law enforcement to other data, as well as the improvement (clearing) or decline (clouding) of said Psycho-Pass. When a targeted individual''s Crime Coefficient index exceeds the accepted threshold (100), they are pursued, apprehended, and either arrested or decomposed by the field officers of the Crime Investigation Department of the Ministry of Welfare''s Public Safety Bureau. Elite officers known as Inspectors research and evaluate crime scenes, including all personnel involved, with the assistance of Enforcers. Enforcers are latent criminals charged with protecting the Inspectors, adding their expertise and carrying out Inspectors'' instructions. Both are equipped with personally activated, hand-held weapons called "Dominators" whose integrated scanners provide the target''s immediate Psycho-Pass. The gun-like weapon can only fire when approved by the Sibyl System and triggered by its owner. Inspectors and Enforcers work as a team, though Inspectors have jurisdiction to fire their Dominators on the Enforcers should they pose a danger to the public or the Inspectors themselves.',
        24, 'cyberpunk, crime, psychological thriller'),
       ('Sinners of the System', '2019-01-25',
        'The first film, Case.1 Tsumi to Batsu (Crime and Punishment), focuses on Nobuchika Ginoza and Mika Shimotsuki as they investigate the prison filled with people living in cheerful life. The second film, Case.2 First Guardian, focuses on Tomomi Masaoka and Teppei Sugo as how they met following the death of Sugo''s friends. The third and final film, Case.3 Onshuu no Kanata ni (On the Other Side of Love and Hate) focuses on Shinya Kogami as he continues his life as a mercenary in Eastern countries while training a young girl seeking for revenge.',
        3, 'cyberpunk, crime'),
       ('First Inspector', '2020-04-27',
        'The film acts as sequel to the 2019 anime Psycho-Pass 3, the third season of Psycho-Pass series. It stars the talents of Yuuki Kaji, Yūichi Nakamura, Mamoru Miyano, Kenyu Horiuchi, among others. Set in a dystopia known as the Sybil System, the film explores the Inspectors Kei Mikhail Ignatov, Arata Shindo, among others who clash with the terrorist group Bifrost in a clash to take Tokyo governor Karina Komiya. The film released to theatres in Japan on March 27, 2020.',
        3, 'crime, psychological thriller'),
       ('Kuroshitsuji', '2008-10-3',
        'In Victorian-era London lives a thirteen-year-old earl named Ciel Phantomhive, who acquired this position after the events of December 14, 1885 (Ciel''s 10th birthday), when the Phantomhive manor was attacked by unidentified perpetrators and was set ablaze. Ciel, amidst the chaos, discovers his parents, Vincent and Rachel Phantomhive, to be dead along with the family dog, Sebastian. The same night he is kidnapped by the attackers and then sold into slavery, where he ends up in the hands of a sadistic, demon-worshipping cult. Ciel then endures endless physical, mental and sexual abuse at the hands of his captors. Ciel was also heat branded with a mark referred to as the "mark of the beast". One night, during a sacrificial ceremony to summon a demon, instead of forming a contract with the cult members, the demon states that he was summoned by Ciel, therefore he only agrees to form a contract with him, killing all the cultist members in the process. To show a contract was formed, the demon places a contract symbol referred to as the Faustian contract symbol on Ciel''s right eye, giving it a purple hue, and his iris and pupil now showcase the symbol of the covenant. After the formation of the contract, the demon reveals he will consume Ciel''s soul as payment for helping him achieve his goal; revenge on those who brought down the House of Phantomhive. Afterwards Ciel names the demon Sebastian Michaelis, after his deceased pet dog. The duo then return to society as Ciel takes over his now late father''s previous position as the queen''s watchdog, a very high-profile individual who is tasked with investigating cases that Queen Victoria herself deems especially important or threatening to England and the crown.',
        24, 'supernatural, comedy, fantasy'),
       ('Book of Circus', '2014-07-10',
        'In late-nineteenth-century England, the Noah''s Ark Circus enlivens the city of London with their grandeur and spectacular stunts. However, children are mysteriously disappearing from the town in a manner correlating to the troupe''s movement. The Queen, then, sends her notorious Watchdog, Ciel Phantomhive, on an investigative mission to retrieve the missing children. Ciel and his demon butler, Sebastian Michaelis, infiltrate the circus, disguised as team members, to study and possibly unearth its ulterior motive.',
        10, 'supernatural, fantasy, crime');

INSERT INTO person (name, birthdate, bio, awards)
VALUES ('Kana Hanazawa', '1989-02-25',
        'She is best known for her voice performances in anime, which include Nadeko Sengoku in Monogatari, Anri Sonohara in Durarara!!, Angel / Kanade Tachibana in Angel Beats!, Kuroneko / Ruri Gokō in Oreimo, Mayuri Shiina in Steins;Gate, Akane Tsunemori in Psycho-Pass, Kosaki Onodera in Nisekoi, Chiaki Nanami in Danganronpa, Kobato Hanato in Kobato, Rize Kamishiro in Tokyo Ghoul, Hinata Kawamoto in March Comes in Like a Lion, Ichika Nakano in The Quintessential Quintuplets and Mitsuri Kanroji in Demon Slayer: Kimetsu no Yaiba. Her debut single, "Hoshizora Destination" (星空☆ディスティネーション), was released on April 25, 2012 under the Aniplex/Sony Music Entertainment Japan label.',
        '9th Seiyu Awards - Best Supporting Actress, Newtype Anime Awards - Best Voice Actress'),
       ('Tomokazu Seki', '1972-09-08',
        'Japanese actor, voice actor, and singer. He has previously worked with Haikyō. He is honorary president of and affiliated with Atomic Monkey and the chairman of theater company HeroHero Q. He is a special lecturer at Japan Newart College.',
        ''),
       ('Kenji Nojima', '1976-04-16',
        'Japanese voice actor and singer affiliated with the voice talent agency Aoni Production. His first major role in voice-over was Spark in the Record of Lodoss War: Chronicles of the Heroic Knight series. He voiced Hikaru Ichijyo in a number of Macross-related video games in the 2000s. Other major roles include Yuto Kiba in High School DxD, Jade in Ultimate Muscle, Nobuchika Ginoza in Psycho-Pass, Keisaku Sato in Shakugan no Shana, Taihei Doma in Himouto! Umaru-chan, and Tuxedo Mask in Sailor Moon Crystal. In anime films, he voices Fumito Nanahara in Blood-C, Masaki in Time of Eve. He is the son of Akio Nojima and is the younger brother of Hirofumi Nojima. He married Chie Sawaguchi in 2004 and has two children.',
        ''),
       ('Miyuki Sawashiro', '1985-06-02',
        'a Japanese actress, voice actress and narrator. She has played voice roles in a number of Japanese anime including Bishamon in Noragami, Petit Charat/Puchiko in Di Gi Charat, Mint in Galaxy Angel, Sinon in Sword Art Online, Twilight/Towa Akagi/Cure Scarlet in Go! Princess Precure, Raiden Mei in Honkai Impact 3, Beelzebub in Beelzebub, Izuna Hatsuse in No Game No Life, Amagi in Azur Lane, Celty Sturluson in Durarara!!, Kurapika in Hunter x Hunter, Raiden Shogun in Genshin Impact, Akane Kurashiki in Zero Escape, Ayane Yano in Kimi ni Todoke, Fujiko Mine in later installments of Lupin the Third, Queen in Mysterious Joker, Jun Sasada in Natsume''s Book of Friends, Shinku in Rozen Maiden, Haruka Nanami in Uta no Prince-sama, Kotoha Isone in Yozakura Quartet, Kanbaru Suruga in Bakemonogatari, Elizabeth and Chidori in Persona 3, Ivy Valentine in Soulcalibur, Jolyne Cujoh in JoJo''s Bizarre Adventure: All Star Battle and JoJo''s Bizarre Adventure: Eyes of Heaven, and Elizabeth in BioShock Infinite.',
        'Seiyu Award - Best Actress in a Leading Role'),
       ('Toshiyuki Morikawa', '1967-01-26',
        'Japanese voice actor, narrator and singer who is the head of Axlone, a voice acting company he founded in April 2011.[1] His name is also sometimes mistranslated as Tomoyuki Morikawa. In 2003, he and Fumihiko Tachiki formed the band "2Hearts", one of their works being the ending theme of the video game Dynasty Warriors 4: Empires. He has voiced many characters in anime and video games, including Yoshikage Kira in JoJo''s Bizarre Adventure: Diamond Is Unbreakable, Kengo Akechi in Kindaichi Case Files, Sephiroth in the Final Fantasy series and Kingdom Hearts series, Dante in Devil May Cry, Kagaya Ubuyashiki in Demon Slayer: Kimetsu no Yaiba, Isshin Kurosaki in Bleach, Minato Namikaze in Naruto: Shippuden, Julius Novachrono in Black Clover, Mard Geer Tartaros in Fairy Tail, Boros in One Punch Man, both Eneru and Hatchan in One Piece, Tyki Mikk in D.Gray-man, Naraku in InuYasha, Griffith in the 1997 series of Berserk, Isaburo Sasaki in Gin Tama and the main and titular character of Tekkaman Blade. He attended Katsuta Voice Actor''s Academy with Kotono Mitsuishi, Chisa Yokoyama, Wataru Takagi, Sachiko Sugawara and Michiko Neya. Because of his deep voice, he is often cast as imposing characters. He has dubbed-over many actors in Japanese such as: Tom Cruise, Ewan McGregor, Adam Sandler, Chris O''Donnell, Owen Wilson, Keanu Reeves, Heath Ledger, Jude Law and Martin Freeman''s roles.',
        'Seiyu Award - Best Actor in a Supporting Role');



INSERT INTO person_show (person_id, show_id)
VALUES ((SELECT person_id FROM person WHERE name = 'Kana Hanazawa'),
        (SELECT show_id FROM show WHERE title = 'Psycho-Pass')),
       ((SELECT person_id FROM person WHERE name = 'Kana Hanazawa'),
        (SELECT show_id FROM show WHERE title = 'Sinners of the System')),
       ((SELECT person_id FROM person WHERE name = 'Kana Hanazawa'),
        (SELECT show_id FROM show WHERE title = 'First Inspector')),
       ((SELECT person_id FROM person WHERE name = 'Kenji Nojima'),
        (SELECT show_id FROM show WHERE title = 'Psycho-Pass')),
       ((SELECT person_id FROM person WHERE name = 'Kenji Nojima'),
        (SELECT show_id FROM show WHERE title = 'Sinners of the System')),
       ((SELECT person_id FROM person WHERE name = 'Kenji Nojima'),
        (SELECT show_id FROM show WHERE title = 'First Inspector')),
       ((SELECT person_id FROM person WHERE name = 'Tomokazu Seki'),
        (SELECT show_id FROM show WHERE title = 'Psycho-Pass')),
       ((SELECT person_id FROM person WHERE name = 'Tomokazu Seki'),
        (SELECT show_id FROM show WHERE title = 'Sinners of the System')),
       ((SELECT person_id FROM person WHERE name = 'Tomokazu Seki'),
        (SELECT show_id FROM show WHERE title = 'First Inspector')),
       ((SELECT person_id FROM person WHERE name = 'Miyuki Sawashiro'),
        (SELECT show_id FROM show WHERE title = 'Psycho-Pass')),
       ((SELECT person_id FROM person WHERE name = 'Miyuki Sawashiro'),
        (SELECT show_id FROM show WHERE title = 'Sinners of the System')),
       ((SELECT person_id FROM person WHERE name = 'Miyuki Sawashiro'),
        (SELECT show_id FROM show WHERE title = 'Book of Circus')),
       ((SELECT person_id FROM person WHERE name = 'Miyuki Sawashiro'),
        (SELECT show_id FROM show WHERE title = 'Kuroshitsuji')),
       ((SELECT person_id FROM person WHERE name = 'Toshiyuki Morikawa'),
        (SELECT show_id FROM show WHERE title = 'Book of Circus')),
       ((SELECT person_id FROM person WHERE name = 'Toshiyuki Morikawa'),
        (SELECT show_id FROM show WHERE title = 'Kuroshitsuji'))
