components {
  id: "flame"
  component: "/main/bots/flame.particlefx"
  position {
    x: 0.0
    y: -24.0
    z: -1.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "jet_script"
  component: "/main/bots/jet_script.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "flame_boosted"
  component: "/main/bots/flame_boosted.particlefx"
  position {
    x: 0.0
    y: -35.0
    z: -1.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "co"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 0.15\n"
  "friction: 0.0\n"
  "restitution: 0.7\n"
  "group: \"jet\"\n"
  "mask: \"wall\"\n"
  "mask: \"body\"\n"
  "mask: \"jet\"\n"
  "mask: \"spikes\"\n"
  "mask: \"bounce\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 0.0\n"
  "      y: -14.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.38268343\n"
  "      w: 0.9238795\n"
  "    }\n"
  "    index: 3\n"
  "    count: 3\n"
  "  }\n"
  "  data: 27.0645\n"
  "  data: 13.5365\n"
  "  data: 10.0\n"
  "  data: 20.0\n"
  "  data: 20.0\n"
  "  data: 10.0\n"
  "}\n"
  "linear_damping: 0.3\n"
  "angular_damping: 0.3\n"
  "locked_rotation: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "jet"
  type: "sprite"
  data: "tile_set: \"/assets/bots/bot1.atlas\"\n"
  "default_animation: \"jet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
